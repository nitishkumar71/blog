FROM alpine:3.9
LABEL author="Nitishkumar Singh"

RUN apk --update --no-cache add python3 python3-dev apache2 apache2-dev  wget ca-certificates make gcc musl-dev;\
ln -s pip3 /usr/bin/pip; pip install -U pip setuptools wheel

# mod_wsgi compilation
RUN wget -O /tmp/mod_wsgi.tar.gz https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/modwsgi/mod_wsgi-3.4.tar.gz && \
    tar -C /tmp -xvf /tmp/mod_wsgi.tar.gz && \
    rm /tmp/mod_wsgi.tar.gz

WORKDIR /tmp/mod_wsgi-3.4
RUN ln -s /usr/lib/libpython3.6m.so /usr/lib/libpython3.6.so && \
    ./configure --with-python=/usr/bin/python3.6 --with-apxs=/usr/bin/apxs && \
    make && make install clean; rm -rf /tmp/mod_wsgi-3.4 \
mkdir -p /var/www/my_django_app; mkdir -p /etc/apache2/sites-available/; \
mkdir -p /etc/apache2/mods-available/

WORKDIR /var/www/my_django_app
COPY . /var/www/my_django_app

# ARG PORT

RUN echo -e "import os\n\
import sys\n\
path='/var/www/my_django_app'\n\
if path not in sys.path:\n\
    sys.path.append(path)\n\
os.environ['DJANGO_SETTINGS_MODULE'] = 'my_django_app.settings'\n\
from django.core.wsgi import get_wsgi_application\n\
application = get_wsgi_application()" >> /var/www/my_django_app/django.wsgi; \
sed -i -r 's@#(LoadModule rewrite_module modules/mod_rewrite.so)@\1@i' /etc/apache2/httpd.conf; \
sed -i -r 's@Errorlog .*@Errorlog /var/log/apache2/error.log@i' /etc/apache2/httpd.conf; \
sed -i -r 's@#Servername .*@ServerName localhost@i' /etc/apache2/httpd.conf; \
sed -i -r 's@Listen 80.*@Listen 8080@i' /etc/apache2/httpd.conf; \
echo -e 'Transferlog /dev/stdout\n\
LoadModule wsgi_module modules/mod_wsgi.so\n\
WSGIPythonPath /usr/lib/python3.6\n\
WSGIScriptAlias / /var/www/my_django_app/django.wsgi\n\
<Directory /var/www/my_django_app>\n\
    Options ExecCGI Indexes FollowSymLinks\n\
    AllowOverride All\n\
    Require all granted\n\
    <Files django.wsgi>\n\
        Require all granted\n\
    </Files>\n\
</Directory>' >> /etc/apache2/httpd.conf
RUN pip install -r requirements.txt; chown apache:apache /etc/apache2/; \
chown -R apache:apache /var/www/
COPY httpd-foreground /usr/local/bin/
EXPOSE 8080
# ENTRYPOINT ["httpd", "-D", "FOREGROUND", "-e", "info", "&"]
CMD ["httpd-foreground"]