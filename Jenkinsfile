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
	stage ('remove stale data'){
	   steps {		
		catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE'){
                    sh  'sshpass -p ${ServerPW} killall jenkinstest'
                    sh  'sshpass -p ${ServerPW} yes | rm /home/jenkinstest' 
		}
	   }
	}
  	stage('deploy'){
	   steps {
		echo 'Deploying ..'
		sh 'mv /root/.jenkins/workspace/jenkinstest_master /root/.jenkins/workspace/jenkinstest'
                sh  'sshpass -p ${ServerPW} scp jenkinstest root@172.19.2.121:/home/'
                sh  'sshpass -p ${ServerPW} ssh root@172.19.2.121 /home/jenkinstest &'
	   }
	}
    }
}
