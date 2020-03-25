# RPSSL-game-server
Rock Paper Scissors Lizard Spock!! Microservice API to run your own RPSSL game

Modifying the Ports to push out HTTP (credit to https://gist.github.com/kentbrew/776580)
Add a port forwarding rule via iptables.

listed the rules currently running on the NAT (Network Address Translation) table:

[ec2-user@ip-XX-XXX-XX-X ~]$ sudo iptables -t nat -L

Chain INPUT (policy ACCEPT)
target     prot opt source    destination

Chain FORWARD (policy ACCEPT)
target     prot opt source    destination

Chain OUTPUT (policy ACCEPT)
target     prot opt source    destination

If nothing, feel free to add a rule forwarding packets sent to external port 80 to internal port 8080:

[ec2-user@ip-XX-XXX-XX-X ~]$ sudo iptables -t nat -A PREROUTING -p tcp --dport 80 -j REDIRECT --to-ports 8080

When listed again, a new PREROUTING chain appears:

[ec2-user@ip-XX-XXX-XX-X ~]$ sudo iptables -t nat -L

Chain PREROUTING (policy ACCEPT)
target     prot opt source     destination
REDIRECT   tcp  --  anywhere   anywhere     tcp dpt:http redir ports 8080

script is now running on port 8080, and it should respond on port 80.