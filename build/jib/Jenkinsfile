pipeline {
  agent {
    kubernetes {
      //cloud 'kubernetes'
      yaml '''
kind: Pod
metadata:
  name: jib
spec:
  containers:
  - name: jnlp
    image: index.alauda.cn/alaudaorg/jnlp-slave:latest
    imagePullPolicy: Always
    workingDir: /home/jenkins/agent
    env:
    - name: JENKINS_URL
      value: http://jenkins.default:8080
    - name: "JENKINS_TUNNEL"
      value: "jenkins-agent.default:50000"
    volumeMounts:
      - mountPath: "/home/jenkins/agent"
        name: "workspace-volume"
        readOnly: false
  - name: jib
    image: index.alauda.cn/alaudak8s/jib
    imagePullPolicy: Always
    command:
    - cat
    tty: true
    workingDir: "/home/jenkins/agent"
    volumeMounts:
    - mountPath: "/home/jenkins/agent"
      name: "workspace-volume"
      readOnly: false
    - mountPath: "/home/root/.m2/repository"
      name: "volume-2"
      readOnly: false
    - mountPath: "/root/maven/repository"
      name: "volume-0"
      readOnly: false
    - mountPath: "/home/root/maven/repository"
      name: "volume-1"
      readOnly: false
    - mountPath: "/var/run/docker.sock"
      name: "docker"
      readOnly: false
  volumes:
  - hostPath:
      path: "/data/mvn1"
    name: "volume-0"
  - hostPath:
      path: "/data/mvn2"
    name: "volume-2"
  - hostPath:
      path: "/data/mvn3"
    name: "volume-1"
  - emptyDir:
      medium: ""
    name: "workspace-volume"
  - hostPath:
      path: "/var/run/docker.sock"
    name: "docker"
'''
    }
  }
  stages{
        stage("Clone"){
           steps{
             git url:"https://github.com/danielfbm/jib-example.git"
           }
        }
        stage("Build"){
        //KO_DOCKER_REPO = "jfrog-demo.alauda.cn/docker"
            environment {
                DOCKER_REPO = "jfrog-demo.alauda.cn/docker/jib-example"
            }
            steps {
            script {
              container('jib') {

                withCredentials([usernamePassword(credentialsId: 'daniel-jfrog', passwordVariable: 'PASS', usernameVariable: 'USER')]) {
                    def authBase = sh script:"echo ${USER}:${PASS} | base64", returnStdout: true
                    sh "docker login $DOCKER_REPO -u ${USER} -p ${PASS}"
                }
                
                sh script: '''
                mkdir -p $HOME/.m2/
cat << EOF > $HOME/.m2/settings.xml
<settings>
<mirrors>
    <mirror>
        <id>nexus-osc</id>
        <mirrorOf>*</mirrorOf>
        <name>Nexus osc</name>
        <url>http://maven.aliyun.com/nexus/content/groups/public/</url>
    </mirror>
</mirrors>
</settings>
EOF
mkdir -p $M2_HOME/conf
cp $HOME/.m2/settings.xml $M2_HOME/conf/settings.xml
'''
                sh script: "mvn compile jib:build"
              }
            }
          }
        }
     }
}