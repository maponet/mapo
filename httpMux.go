package main

import (
    "mapo/log"
    
    "os"
    "sync"
    "time"
    "regexp"
    "strings"
    "net/http"
)

// handler, personalizzato per il server http che ci permetterà di spegnere
// l'applicazione senza rischi o corruzione dei dati.
type ServeMux struct {

    // lista dei handler registrati
    m map[string]Handler
    
    authenticator func(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request, bool)

    // il numero delle connessione attive in questo momento
    current_connections int
    lock sync.Mutex
    
    // il server è o no in fase di chiusura
    closing bool
}

type handlerWithAuthentication struct {
    f func(http.ResponseWriter, *http.Request)
    auth func(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request, bool)
}
func (hwa handlerWithAuthentication) ServeHTTP(out http.ResponseWriter, in *http.Request) {
    if hwa.auth != nil {
        if o, i, ok := hwa.auth(out, in); ok {
            hwa.f(o, i)
        }
    } else {
        hwa.f(out, in)
    }
}

type handlerWithoutAuthentication func(http.ResponseWriter, *http.Request)
func (hwoa handlerWithoutAuthentication) ServeHTTP(out http.ResponseWriter, in *http.Request) {
    hwoa(out, in)
}

func (mux *ServeMux) SetAuthenticator(auth func(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request, bool)) {
    mux.authenticator = auth
}

func (mux *ServeMux) HandleFunc(method, path string, handle func(http.ResponseWriter, *http.Request) ) {

    handlerFunc := new(handlerWithAuthentication)
    handlerFunc.auth = mux.authenticator
    handlerFunc.f = handle

    pattern := createPattern(method, path)
    
    mux.m[pattern] = handlerFunc
}

func (mux *ServeMux) HandleFuncNoAuth(method, path string, handle func(http.ResponseWriter, *http.Request) ) {

    handlerFunc := new(handlerWithoutAuthentication)
    *handlerFunc = handle

    pattern := createPattern(method, path)
    
    mux.m[pattern] = handlerFunc
}

func createPattern(method, path string) string {
    pattern := "(?i)^"
    
    if method != "" {
        pattern = pattern + method + ":/"
    } else {
        pattern = pattern + "(GET|POST)" + ":/"
    }
    
    if len(path) > 1 {
        pathSlice := strings.Split(path[1:], "/")
        for _, v := range(pathSlice) {
            if v[0] == '{' {
                pattern = pattern + "[0-9a-z_\\ \\.\\+\\-]*/"
            } else {
                pattern = pattern + v + "/"
            }
        }
    }
    pattern = pattern + "$"
    
    return pattern
}

func (mux *ServeMux) match(r *http.Request) Handler {
    method := r.Method
    url := r.URL.Path
    
    if url[len(url)-1] != '/' {
        url = url + "/"
    }
    
    var handler Handler
    
    for k, v := range(mux.m) {
        matching, _ := regexp.MatchString(k, method + ":" + url)
        if matching {
            handler = v
            break
        }
    }
    
    if handler != nil {
        return handler
    }
    return http.NotFoundHandler()
}

func NewServeMux() *ServeMux {
    mux := new(ServeMux)
    mux.m = make(map[string]Handler, 0)
    
    return mux
}

type Handler interface {
    ServeHTTP(http.ResponseWriter, *http.Request)
}

// ServeHTTP e la funzione che vine eseguita come gorutine ogni volta che
// si deve processare qualche richiesta. Questa funzione soltanto si assicura
// che venga incrementato o decrementato il numero delle connessione attive e
// avvierà la funzione RequestHandler che processerà la richiesta del cliente.
// Comunque, il server http viene interrotto in maniera brutta ma senza alcun
// rischio. TODO: approfondire questa feature se servirà.
func (mux *ServeMux) ServeHTTP(out http.ResponseWriter, in *http.Request) {
    if !mux.closing {
        start := time.Now()
        defer func() {
            log.Info("time: %v for %s", time.Since(start), in.URL.Path)
        }()
        
        mux.lock.Lock()
        mux.current_connections++
        mux.lock.Unlock()
        
        defer func() {
            mux.lock.Lock()
            mux.current_connections--
            mux.lock.Unlock()
        }()
        
        handle := mux.match(in)
        handle.ServeHTTP(out, in)
    }
}

// se viene richiesto che l'applicazione si deve chiudere, in questo momento si
// parla del commando CTRL+C dal terminale, potremmo corrompere i dati a colpa
// del'interruzione in maniera incorretta delle richieste in corso. La presente
// Funzione sta in ascolto per il segnale SIGINT dopo di che si assicura che il
// server venga chiuso non appena le connessione attive saranno zero.
func (mux *ServeMux) getSignalAndClose(c chan os.Signal) {

    _ = <-c
    log.Info("closing ...")
    mux.closing = true
    
    // TODO: send notification to load balancing that this node is unavailable
    
    for {
        if mux.current_connections == 0 {
            log.Info("bye ... :)")
            os.Exit(1)
        } else {
            log.Info("waiting for %d opened connections", mux.current_connections)
            time.Sleep(500 * time.Millisecond)
        }
    }
}