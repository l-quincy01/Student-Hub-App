## Overview
Student Hub is a community marketplace for college and university students to buy and sell second-hand goods, fostering a sustainable and affordable exchange of textbooks, furniture, electronics, and other essentials within the student community.

## User Interface
#### Onboarding Screens
| Welcome  | Login | Sign Up |
|--------|--------|--------|
| <img src="./Screenshots/LoginScreen.PNG" width="85%" /> | <img src="./Screenshots/LoginScreen.PNG" width="85%" /> | <img src="./Screenshots/LoginScreen.PNG" width="85%" /> | 

## Folder Structure & Tech Stack
The folder structructe follows the convetions of Go & React Native applications as depicted below:
```
├── api
│   ├── cmd
│   ├── config
│   ├── db
│   └── internal
└── client
    ├── App.js
    ├── Apps
    ├── assets
    ├── hooks
    ├── package.json
    └── tsconfig.json
```
The tech stack is made up of React Native & Golang as depicted below:

| **Language**  | **Framework** |
|--------|--------|
| Javascript | React Native |


| **Language**  | **Framework** |
|--------|--------|
| Go | Fiber |

|  **Database** | **BaaS** |
|--------|--------|
| MongoDB | Firebase | 


## Authentication & Authorization
The Student-Hub will be using Firebase as a provider for handling OAuth 2.0 and phone number authetication. Below is a high level overview of the flow between the client, API & Firebase:

```mermaid
flowchart LR
    id0[React App] -- authenitcation  --> id1[[FirebaseAuth]]
    id2[[API]] -- manages user info --> id1 
    id2 <-- update or create users --> id3[(MongoDB)]
    id0 -- requests --> id2
```

#### Sign in/Sign up Sequence Diagram
```mermaid
sequenceDiagram
   actor  Client
   Client ->>FirebaseAuth: signs up/signs in with oauth2 or phone number 
    alt is phone number login
        FirebaseAuth ->> Client: send OTP
        Client ->>FirebaseAuth: verify OTP
        alt OTP invalid
            FirebaseAuth->>Client: notify invalid OTP
        end
    end
    FirebaseAuth ->> Client: authenticate user

    Client->>API: request account info
    alt is new user
        API ->> FirebaseAuth: get signed-in user info
        API ->> MongoDB: create user account
        API->> Client: notify user account is new
        Client->>API: send additional user info
        API->>MongoDB: update user account
    else is existing user
        API->>MongoDB: fetch account information
    end
    API ->> Client: send account info
```