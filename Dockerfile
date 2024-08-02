FROM golang:1.19-alpine

LABEL maintainer="coko@duck.com"

###############################################################################
#                                INSTALLATION
###############################################################################
# Set project path
ENV WORKDIR /var/www/hertz-admin
# Add the application executable and set the execution permission
ADD ./hertz-admin   $WORKDIR/hertz-admin
RUN chmod +x $WORKDIR/hertz-admin

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
# Set the environment variables
ENV IS_PROD true
# Run the application
CMD ./hertz-admin
