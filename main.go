package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

const (
	// PublicNtpServer определяет адрес публичного NTP-сервера.
	PublicNtpServer = "pool.ntp.org"
)

func main() {
	// Используем функцию Time() для простого получения уже скорректированного времени.
	// Эта функция берет на себя всю логику запроса и расчета смещения.
	exactTime, err := ntp.Time(PublicNtpServer)

	if err != nil {
		// В случае ошибки выводим сообщение в STDERR и завершаем программу с ненулевым кодом выхода.
		log.Fatal("Ошибка при запросе времени NTP:", err)
	}

	// Используем стандартный формат RFC3339 для читаемости.
	fmt.Printf("Точное текущее время (NTP): %s\n", exactTime.Format(time.RFC3339))

}
