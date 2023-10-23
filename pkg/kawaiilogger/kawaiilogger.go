package kawaiilogger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/forgetz/go-kawaii-shop/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type IKawaiiLogger interface {
	Print() IKawaiiLogger
	Save()
	SetQuery(c *fiber.Ctx)
	SetBody(c *fiber.Ctx)
	SetResponse(res any)
}

type kawaiilogger struct {
	Time       string `json:"time"`
	Ip         string `json:"ip"`
	Method     string `json:"method"`
	StatusCode int    `json:"status_code"`
	Path       string `json:"path"`
	Query      any    `json:"query"`
	Body       any    `json:"body"`
	Response   any    `json:"response"`
}

func InitKawaiiLogger(c *fiber.Ctx, res any) IKawaiiLogger {
	log := &kawaiilogger{
		Time:       time.Now().Local().Format("2006-01-02 15:04:05"),
		Ip:         c.IP(),
		Method:     c.Method(),
		StatusCode: c.Response().StatusCode(),
		Path:       c.Path(),
	}

	log.SetQuery(c)
	log.SetBody(c)
	log.SetResponse(res)
	return log
}

func (l *kawaiilogger) Print() IKawaiiLogger {
	utils.Debug(l)
	return l
}

func (l *kawaiilogger) Save() {
	data := utils.Output(l)
	filename := fmt.Sprintf("./assets/logs/kawaiilogger_%v.txt", strings.ReplaceAll(time.Now().Format("2006-01-02"), "-", ""))
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	file.WriteString(string(data) + "\n")
}

func (l *kawaiilogger) SetQuery(c *fiber.Ctx) {
	var body any
	if err := c.QueryParser(&body); err != nil {
		log.Fatalf("error query parser: %v", err)
	}
	l.Query = body
}

func (l *kawaiilogger) SetBody(c *fiber.Ctx) {
	var body any
	if err := c.BodyParser(&body); err != nil {
		log.Fatalf("error body parser: %v", err)
	}

	switch l.Path {
	case "v1/users/signup":
		l.Body = "never gonna give you up"
	default:
		l.Body = body
	}
}

func (l *kawaiilogger) SetResponse(res any) {
	l.Response = res
}
