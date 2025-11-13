pipeline {
    environment {
        QODANA_TOKEN = credentials('qodana-token')
    }
    agent {
        docker {
            args '''
                -v "${WORKSPACE}":/data/project
                --entrypoint=""
                '''
            image 'jetbrains/qodana-go'
        }
    }
    stages {
        stage('Qodana') {
            when {
                branch 'dev'
                branch 'main'
            }
            steps {
                sh '''qodana --baseline=qodana.sarif.json'''
            }
        }
    }
}