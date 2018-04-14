A program is simple app that help to send and receive the messages to/from message queue

user UI → app → Queue(kafka) → app → user UI

<b>Notes</b>:
- This app will be deployed using Docker one container per module.
- You can write in any programing language.
- Write Ansible modules to build and deploy this app to server (you can use your local host).

<b>Screen</b>

Sender:</p>
<img src="https://raw.githubusercontent.com/bienkma/SentAndRecivedMsgQueue/master/screen/sender.png"/>

Receiver:</p>
<img src="https://raw.githubusercontent.com/bienkma/SentAndRecivedMsgQueue/master/screen/reciever.png"/>
