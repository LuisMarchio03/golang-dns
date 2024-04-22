# Modelo OSI (Open Systems Interconnection)

1. **Camada Física**
   - Esta camada lida com a transmissão física dos dados sobre o meio de comunicação. Isso inclui a definição de características elétricas, mecânicas e funcionais do hardware, como cabos, conectores e sinais elétricos.

2. **Camada de Enlace de Dados**
   - Responsável pelo estabelecimento, manutenção e término da conexão física entre dispositivos. Ela divide os dados em quadros e adiciona informações de controle de fluxo e correção de erros. Garante a entrega confiável dos dados no nível local.

3. **Camada de Rede**
   - Gerencia o roteamento dos dados entre redes diferentes. Ela encapsula os dados em pacotes, adicionando endereçamento e controle de roteamento. É onde os endereços IP são utilizados para determinar o caminho dos pacotes na rede.

4. **Camada de Transporte**
   - Oferece mecanismos para transferência de dados confiável e eficiente entre hosts. O TCP (Transmission Control Protocol) opera nesta camada, garantindo a entrega ordenada e sem erros dos dados. Além disso, esta camada lida com o controle de congestionamento e a segmentação dos dados.

5. **Camada de Sessão**
   - Gerencia a comunicação entre processos em diferentes sistemas. Estabelece, mantém e finaliza sessões de comunicação entre aplicativos, incluindo a sincronização e recuperação de dados em caso de falhas.

6. **Camada de Apresentação**
   - Responsável pela tradução, compressão e criptografia dos dados para a comunicação entre sistemas heterogêneos. Garante que os dados sejam apresentados de forma correta para o aplicativo, independentemente do formato em que foram transmitidos pela rede.

7. **Camada de Aplicação**
   - Fornece interfaces para os aplicativos de usuário acessarem os serviços de rede. Inclui protocolos como HTTP, FTP, SMTP, entre outros, que permitem a comunicação e interação entre aplicativos em diferentes computadores.

# Modelo TCP/IP (Transmission Control Protocol/Internet Protocol)

1. **Camada de Aplicação**
   - Fornece serviços de rede diretamente aos aplicativos. Inclui protocolos como HTTP, FTP, SMTP, DNS, entre outros, que são usados para interação entre aplicativos.

2. **Camada de Transporte**
   - Responsável pela comunicação ponto a ponto entre hosts. O TCP e o UDP (User Datagram Protocol) são os principais protocolos nesta camada. O TCP fornece comunicação confiável, enquanto o UDP é usado para comunicação não confiável, porém mais rápida.

3. **Camada de Internet**
   - Gerencia o roteamento dos dados entre redes diferentes. O protocolo IP (Internet Protocol) opera nesta camada, fornecendo endereçamento e roteamento dos pacotes pela Internet.

4. **Camada de Interface de Rede**
   - Responsável pela comunicação direta com o hardware de rede. Ela lida com a transmissão e recepção dos dados nos meios físicos, como cabos de rede e interfaces de rede.

Estas camadas formam uma estrutura hierárquica que permite a comunicação eficiente e confiável entre dispositivos em uma rede de computadores. Cada camada possui suas próprias funções e responsabilidades, contribuindo para a comunicação global na rede.
