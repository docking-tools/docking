FROM centos:6
MAINTAINER Sebastien GLON <glon.sebastien@free.fr>

ADD sensu.repo /etc/yum.repos.d/sensu.repo

RUN yum update -y && rpm -Uvh http://dl.fedoraproject.org/pub/epel/6/i386/epel-release-6-8.noarch.rpm \
 && yum install -y sensu \
 && /opt/sensu/embedded/bin/gem install sensu-plugins-docker

ADD config.json /etc/sensu/config.json
#ADD client.json /etc/sensu/conf.d/client.json
ADD docker-check.json /etc/sensu/conf.d/docker-check.json 
RUN chown -R sensu:sensu /etc/sensu


EXPOSE 3030

ADD docker-entrypoint.sh /docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]

CMD /opt/sensu/bin/sensu-client -c /etc/sensu/config.json -d /etc/sensu/conf.d start




