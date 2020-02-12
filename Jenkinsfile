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
                    sh  'sshpass -p ${ServerPW} ssh root@172.19.2.121 killall jenkinstest'
		}
		catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE'){
                    sh  'sshpass -p ${ServerPW} ssh root@172.19.2.121 "yes | rm /home/jenkinstest"' 
		}

	   }
	}
  	stage('deploy'){
	   steps {
		echo 'Deploying ..'
		sh 'mv /root/.jenkins/workspace/jenkinstest_master/jenkinstest_master /root/.jenkins/workspace/jenkinstest_master/jenkinstest'
                sh  'sshpass -p ${ServerPW} scp jenkinstest root@172.19.2.121:/home/'
                sh  'sshpass -p ${ServerPW} ssh root@172.19.2.121 /home/jenkinstest &'
	   }
	}
    }
}
