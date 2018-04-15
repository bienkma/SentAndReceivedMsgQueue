A program is simple app that help to send and receive the messages to/from message queue

user UI → app → Queue(kafka) → app → user UI

<b>Notes</b>:
- This app will be deployed using Docker one container per module.
- You can write in any programing language.
- Write Ansible modules to build and deploy this app to server (you can use your local host).

<b>Install:</b>
- Edit IP_HOST for app binding
```apple js
$ git clone https://github.com/bienkma/SentAndReceivedMsgQueue.git
$ cd SentAndReceivedMsgQueue
$ vi .env 
$ vi ansible.yml
  vars:
    IP_HOST: "192.168.1.47" # changeme
```
- Install with docker-compose
```apple js
$ cd SentAndReceivedMsgQueue
$ cd app; sudo docker build -t app .
$ cd ../ui; sudo docker build -t ui .
$ sudo docker-compose -f docker-compose.yml up -d
$ sudo docker ps
```
- Install with ansible

```apple js
$ sudo apt-get update
$ sudo apt-get install software-properties-common
$ sudo apt-add-repository ppa:ansible/ansible
$ sudo apt-get update
$ sudo apt-get install ansible
$ ansible-playbook ansible-playbook docker-deploy.yml
```

<b>Run:</b>
```apple js
$ sudo  docker exec -it 42d7d669193e /bin/bash
$ ./ui
s enter for send msg
clt-c quit

$./ui
r enter for receive msg
clt-c quit
```

<b>Screen:</b>
- Sender & Receiver
<img src="https://raw.githubusercontent.com/bienkma/SentAndReceivedMsgQueue/master/screen/enter_ui.png"/>
<img src="https://raw.githubusercontent.com/bienkma/SentAndReceivedMsgQueue/master/screen/sender_receiver.png"/>
