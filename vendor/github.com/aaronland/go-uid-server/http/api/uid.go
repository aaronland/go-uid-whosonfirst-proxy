package api

import (
	"fmt"
	"net/http"

	"github.com/aaronland/go-uid"
	"github.com/sfomuseum/go-http-auth"		
)

type UIDHandlerOptions struct {
	Provider uid.Provider
	Authenticator auth.Authenticator
}

func UIDHandler(opts *UIDHandlerOptions) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()
		
		id, err := opts.Provider.UID(ctx)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		str_id := fmt.Sprintf("%v", id.Value())
			
		rsp.Header().Set("Content-type", "text/plain")
		rsp.Write([]byte(str_id))
	}

	return http.HandlerFunc(fn), nil	
}
