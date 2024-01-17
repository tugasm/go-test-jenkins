pipeline {
    agent any

    environment {
        registry = "tugasm1998" 
        imageName = "first-go-app"
        imageTag = "latest"
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
                    docker.withRegistry('https://registry.hub.docker.com', '1cf785e3-b19b-45a6-890e-b178c3f57ea6') {
              
                        def customImage = docker.build("${registry}/${imageName}:${imageTag}")

                        customImage.push()
                    }
                }
            }
        }
    }
}
