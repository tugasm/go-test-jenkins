pipeline {
    agent any

    environment {
        registry = "tugasm1998" 
        imageName = "first-go-app"
        imageTag = "latest"
    }

    tools {
        docker 'docker-tool'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build and Push Docker Image') {
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'tugasm-dockerhub') {
                        docker.build("${registry}/${imageName}:${imageTag}").push()
                    }
                }
            }
        }
    }
}
