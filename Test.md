Testar servidor

go run main.go

nslookup meubanco.local -port=8053

Para obter o IP do seu banco de dados MySQL rodando em um contêiner Docker, você pode usar o seguinte comando:

docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' nome_do_contêiner_mysql

Config BIND:

1. Instalação do BIND:
bash
Copy code
sudo apt update
sudo apt install bind9
2. Edição do arquivo de configuração do BIND:
bash
Copy code
sudo nano /etc/bind/named.conf.local
Adicione as seguintes linhas ao final do arquivo para definir uma nova zona:

plaintext
Copy code
zone "meubanco.local" {
    type master;
    file "/etc/bind/db.meubanco.local"; // Caminho para o arquivo de zona
};
Salve e feche o arquivo.

3. Criação do arquivo de zona:
bash
Copy code
sudo nano /etc/bind/db.meubanco.local
Adicione o seguinte conteúdo ao arquivo de zona:

plaintext
Copy code
$TTL 3600
@       IN      SOA     ns1.meubanco.local. admin.meubanco.local. (
                           2024041901      ; Serial
                           3600            ; Refresh
                           1800            ; Retry
                           604800          ; Expire
                           86400 )         ; Minimum TTL

@       IN      NS      ns1.meubanco.local.
@       IN      A       172.29.0.4
Substitua 172.29.0.4 pelo endereço IP correto do seu contêiner Docker.

Salve e feche o arquivo.

4. Verificação da sintaxe e reinício do BIND:
Verifique se a sintaxe do arquivo de configuração está correta:

bash
Copy code
sudo named-checkconf
Se não houver erros, reinicie o serviço BIND:

bash
Copy code
sudo systemctl restart bind9
5. Teste da configuração:
Use o comando nslookup para testar se o DNS está respondendo corretamente às consultas:

bash
Copy code
nslookup meubanco.local
Você deve ver uma resposta com o endereço IP configurado.

Converta tudo isso para um codigo markdown pfv.;.. preciso que vc me desponibilize todas essas infos em codigo markdown
