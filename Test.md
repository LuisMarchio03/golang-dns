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


sudo /usr/local/go/bin/go run teste.go 




dev-02@pop-os:~/IdeaProjects/golang-dns$ systemctl status named.service
● named.service - BIND Domain Name Server
     Loaded: loaded (/lib/systemd/system/named.service; enabled; vendor preset: enabled)
     Active: active (running) since Tue 2024-04-23 07:45:02 -03; 6h ago
       Docs: man:named(8)
    Process: 1055 ExecStart=/usr/sbin/named $OPTIONS (code=exited, status=0/SUCCESS)
   Main PID: 1105 (named)
      Tasks: 62 (limit: 18667)
     Memory: 27.1M
        CPU: 2.351s
     CGroup: /system.slice/named.service
             └─1105 /usr/sbin/named -u bind

abr 23 14:05:17 pop-os named[1105]: network unreachable resolving 'e10856.b.akamaiedge.net/HTTPS/IN': 2600:1480:e800::c0#53
abr 23 14:05:17 pop-os named[1105]: network unreachable resolving 'e10856.b.akamaiedge.net/A/IN': 2600:1480:e800::c0#53
abr 23 14:05:45 pop-os named[1105]: validating ab.chatgpt.com/A: no valid signature found
abr 23 14:05:45 pop-os named[1105]: validating ab.chatgpt.com/HTTPS: no valid signature found
abr 23 14:05:53 pop-os named[1105]: network unreachable resolving 'alive.github.com/A/IN': 2600:9000:5306:ab00::1#53
abr 23 14:05:53 pop-os named[1105]: network unreachable resolving 'alive.github.com/HTTPS/IN': 2600:9000:5306:ab00::1#53
abr 23 14:05:53 pop-os named[1105]: network unreachable resolving 'alive.github.com/A/IN': 2600:9000:5305:300::1#53
abr 23 14:05:53 pop-os named[1105]: network unreachable resolving 'alive.github.com/HTTPS/IN': 2600:9000:5305:300::1#53
abr 23 14:05:53 pop-os named[1105]: network unreachable resolving 'alive.github.com/A/IN': 2600:9000:5302:800::1#53
abr 23 14:05:53 pop-os named[1105]: network unreachable resolving 'alive.github.com/HTTPS/IN': 2600:9000:5302:800::1#53
dev-02@pop-os:~/IdeaProjects/golang-dns$ 



## sudo nano /etc/resolv.conf





dev-02@pop-os:~/IdeaProjects/golang-dns$ ps aux | grep ./main
postgres    1211  0.0  0.0 232820  9120 ?        SNs  07:45   0:00 /usr/lib/postgresql/16/bin/postgres -D /var/lib/postgresql/16/main -c config_file=/etc/postgresql/16/main/postgresql.conf
postgres    1285  0.0  0.0 232952  4704 ?        SNs  07:45   0:00 postgres: 16/main: checkpointer 
postgres    1286  0.0  0.0 232968  3744 ?        SNs  07:45   0:00 postgres: 16/main: background writer 
postgres    1291  0.0  0.0 232820  4064 ?        SNs  07:45   0:00 postgres: 16/main: walwriter 
postgres    1292  0.0  0.0 234420  4704 ?        SNs  07:45   0:00 postgres: 16/main: autovacuum launcher 
postgres    1293  0.0  0.0 234392  4064 ?        SNs  07:45   0:00 postgres: 16/main: logical replication launcher 
dev-02     18915  0.8  0.4 1177313108 71356 ?    Sl   07:47   3:09 /usr/share/code/code --ms-enable-electron-run-as-node /home/dev-02/.vscode/extensions/streetsidesoftware.code-spell-checker-3.0.1/packages/_server/dist/main.cjs --node-ipc --clientProcessId=15845
dev-02     19695  0.9  0.4 1177313108 72636 ?    Sl   07:47   3:47 /usr/share/code/code --ms-enable-electron-run-as-node /home/dev-02/.vscode/extensions/streetsidesoftware.code-spell-checker-3.0.1/packages/_server/dist/main.cjs --node-ipc --clientProcessId=18543
dev-02    716695  0.8  0.8 1177313108 141472 ?   Sl   12:55   0:35 /usr/share/code/code --ms-enable-electron-run-as-node /home/dev-02/.vscode/extensions/streetsidesoftware.code-spell-checker-3.0.1/packages/_server/dist/main.cjs --node-ipc --clientProcessId=716019
root      895627  0.0  0.0  26164  6696 pts/3    SN+  14:02   0:00 sudo ./main
root      895841  0.0  0.0  26164  2584 pts/5    SNs  14:02   0:00 sudo ./main
root      895842  0.0  0.0 1526652 5600 pts/5    SNl+ 14:02   0:00 ./main
dev-02    909289  0.0  0.0  19040  2560 pts/8    SN+  14:08   0:00 grep --color=auto ./main
dev-02@pop-os:~/IdeaProjects/golang-dns$ 
