FROM yiisoftware/yii2-php:7.3-apache

WORKDIR /app
COPY . /app
RUN chown -Rf www-data:www-data /app
EXPOSE 80
