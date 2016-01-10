FROM centos:6
MAINTAINER Sebastien GLON <glon.sebastien@free.fr>

ADD sensu.repo /etc/yum.repos.d/sensu.repo

RUN yum update -y && rpm -Uvh http://dl.fedoraproject.org/pub/epel/6/i386/epel-release-6-8.noarch.rpm \
 && yum install -y sensu supervisor \
 && /opt/sensu/embedded/bin/gem install sensu-plugins-docker

ADD config.json /etc/sensu/config.json
ADD client.json /etc/sensu/conf.d/client.json
ADD supervisord.conf /etc/supervisord.conf
RUN chown -R sensu:sensu /etc/sensu


EXPOSE 3030

CMD ["/usr/bin/supervisord"]




