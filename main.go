package main

import (
	"log"
	"time"

	sysinfo "github.com/Jefschlarski/monitor/src/application/services/sys_info"
	"github.com/Jefschlarski/monitor/src/configs"
	"github.com/Jefschlarski/monitor/src/dto"
	"github.com/Jefschlarski/monitor/src/kafka"
	"github.com/Jefschlarski/monitor/src/repository/repo"
)

func main() {

	err := configs.Load()
	if err != nil {
		log.Fatal("Erro ao tentar carregar as configurações", err)
	}

	log.Print("Configurações carregadas com sucesso")

	kafkaConfigs := configs.GetKafkaConfig()

	p, err := kafka.NewProducer(kafkaConfigs.HostPort())
	if err != nil {
		log.Fatal("Erro ao tentar criar um novo producer kafka", err)
	}
	defer p.Close()

	log.Printf("Producer conectado com sucesso ao broker Kafka %s", kafkaConfigs.HostPort())

	repo := repo.NewSystemInfoRepo()

	cpuUsageRepo := sysinfo.NewGetCpuUsage(repo)
	diskUsageRepo := sysinfo.NewGetDiskUsage(repo)
	memoryUsageRepo := sysinfo.NewGetMemoryUsage(repo)
	memorySwapUsageRepo := sysinfo.NewGetMemorySwapUsage(repo)

	interval := 1 * time.Second

	ticker := time.NewTicker(interval)

	for range ticker.C {
		cpuUsage, err := cpuUsageRepo.GetCpuUsage()
		if err != nil {
			log.Print("Erro ao obter uso da CPU:", err)
		}

		diskUsage, err := diskUsageRepo.GetDiskUsage()
		if err != nil {
			log.Print("Erro ao obter uso do disco:", err)
		}

		memoryUsage, err := memoryUsageRepo.GetMemoryUsage()
		if err != nil {
			log.Print("Erro ao obter uso da memória:", err)
		}

		swapUsage, err := memorySwapUsageRepo.GetMemorySwapUsage()
		if err != nil {
			log.Print("Erro ao obter uso da swap:", err)
		}

		infoDto := dto.NewInfoDto(memoryUsage, cpuUsage, diskUsage, swapUsage)

		message := kafka.NewMessage(infoDto, kafkaConfigs.Topic())
		err = p.Produce(message, nil)
		if err != nil {
			log.Print("Erro ao enviar mensagem:", err)
		}

		log.Printf("Mensagem enviada com sucesso ao topico: %s", kafkaConfigs.Topic())

		p.Flush(500)
	}
}
