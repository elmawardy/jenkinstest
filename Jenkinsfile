pipeline {
    agent any
    environment {
        ServerPW    = credentials('server-pw')
    }
    stages {
        stage('build') {
            steps {
                sh 'go build'
            }
        }
	stage('test'){
	   steps {
	        sh 'go test'
	   }
	}
  	stage('deploy'){
	   steps {
		echo 'Deploying ..'
                sh  'sshpass -p ${ServerPW} scp jenkinstest root@172.19.2.100:/home/'
                sh  'sshpass -p ${ServerPW} ssh root@172.19.2.100 /home/jenkinstest &'
	   }
	}
    }
}
