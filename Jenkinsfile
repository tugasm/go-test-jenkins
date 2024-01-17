pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Build & Dockerize') {
            steps {
                script{
                    docker.build("first-go-app")
                }
            }
        }
        stage('Push to Docker Hub') {
            steps {
                script{
                    docker.withRegistry('https://registry.hub.docker.com', '1cf785e3-b19b-45a6-890e-b178c3f57ea6') {
                        docker.build("tugasm1998/first-go-app")
                        docker.image("tugasm1998/first-go-app").push()
                    }
                }
            }
        }
    }
}