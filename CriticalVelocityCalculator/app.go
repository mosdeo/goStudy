package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// from uuid import uuid4
	uuid "github.com/satori/go.uuid"
)

// from flask import Flask, request, render_template
// import logging
// import logging.config
// import yaml
// from kh import predict

// app = Flask(__name__, static_url_path='')

// try:
//     with open('logging.yml') as fd:
//         conf = yaml.load(fd)
//         logging.config.dictConfig(conf['logging'])
// except OSError:
//     conf = None

// logger = logging.getLogger('app')
// input_logger = logging.getLogger('app.input')

// if conf:
//     logger.info('logging.yml found, applying config')
//     logger.debug(conf)
// else:
//     logger.info('logging.yml not found')

// def root():
//     uuid = request.cookies.get('uuid', uuid4())
//     resp = app.send_static_file('index.html')
//     resp.set_cookie('uuid', str(uuid))
//     return resp
func root(c *gin.Context) {
	uuidCookie, _ := c.Cookie("uuid")
	c.HTML(http.StatusOK, "index.html", gin.H{})
	c.SetCookie("uuid", uuidCookie, 3600, "/", "localhost", false, true)
}

// def ask():
//     rec = {'ip': ip(),
//            'uuid': request.cookies.get('uuid'),
//            'data': request.form.get('in')}
//     input_logger.info(rec)
//     return predict(request.form.get('in'))
func ask(c *gin.Context) {
	ip := ip(c.Request)
	uuidCookie, err := c.Cookie("uuid")
	if err != nil {
		uuidCookie = "NotSet"
		u1 := uuid.Must(uuid.NewV4())
		c.SetCookie("uuid", u1.String(), 3600, "/", "localhost", false, true)
	}
	data := c.Request.PostFormValue("in")
	fmt.Printf("ip=%v, uuid=%v, data=%v\n", ip, uuidCookie, data)
	c.String(http.StatusOK, "Bot echo=[%v]", data)
}

// def ip():
//     return request.environ.get('REMOTE_ADDR', request.remote_addr)
func ip(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

// if __name__ == '__main__':
func main() {
	router := gin.Default()
	router.StaticFile("/app.js", "./static/app.js")
	router.LoadHTMLFiles("static/index.html")

	// @app.route('/')
	router.GET("/", root)

	// @app.route('/ask', methods=['GET', 'POST'])
	router.GET("/ask", ask)
	router.POST("/ask", ask)

	//     app.run()
	router.Run(":80")
}
