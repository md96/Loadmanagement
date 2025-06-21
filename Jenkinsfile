pipeline {
    agent {
        label 'windows'   // Replace with your Windows Jenkins node label
    }

    environment {
        PATH = "${env.PATH};C:\\Program Files\\Docker\\Docker\\resources\\bin"
    }

    stages {
        stage('Checkout') {
            steps {
                git url: 'https://github.com/md96/Loadmanagement.git', branch: 'main'
            }
        }

        stage('Build Go App') {
            steps {
                bat 'go mod tidy'
                bat 'go build -o load-management.exe .'
            }
        }

        stage('Build Docker Image') {
            steps {
                bat 'docker build -t load-management .'
            }
        }

        stage('Run Docker Container') {
            steps {
                bat 'docker run --rm load-management'
            }
        }
    }

    post {
        always {
            echo 'Pipeline finished.'
        }
    }
}