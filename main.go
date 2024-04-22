package main

import (
	"fmt"
	"os"

	"github.com/miekg/dns"
)

func dnsHandler(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative = true

	for _, q := range r.Question {
		if q.Name == "meubanco.local." && q.Qtype == dns.TypeA {
			rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, "172.29.0.4")) // Substitua "127.0.0.1" pelo endereço IP correto do seu banco de dados
			if err != nil {
				fmt.Println("Erro ao criar resposta DNS:", err)
				return
			}
			m.Answer = append(m.Answer, rr)
		}
	}

	if err := w.WriteMsg(m); err != nil {
		fmt.Println("Erro ao enviar resposta DNS:", err)
	}
}

func main() {
	dns.HandleFunc(".", dnsHandler)
	port := ":8053"
	server := &dns.Server{Addr: port, Net: "udp"}
	fmt.Printf("Servidor DNS iniciado na porta %s\n", port)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Erro ao iniciar o servidor DNS:", err)
		os.Exit(1)
	}
}
