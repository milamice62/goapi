pipeline {
    agent any
    stages {
        stage("Unit Test") {
            agent {
                docker {
                    image 'golang:latest'
                }
            }
            steps {
                echo "Unit Test Stage"
                sh "go version"
            }
        }
        stage("Docker Build") {
            when {
                branch "master"
            }
            failFast true
            parallel {
                stage("Branch A") {
                    steps {
                        echo "On Branch A"
                    }
                }
                stage('Branch B') {
                    steps {
                        echo "On Branch B"
                    }
                }
                stage('Branch C') {
                    stages {
                        stage('Nested 1') {
                            steps {
                                echo "In stage Nested 1 within Branch C"
                            }
                        }
                        stage('Nested 2') {
                            steps {
                                echo "In stage Nested 2 within Branch C"
                            }
                        }
                    }
                }
            }
        }
    }
}