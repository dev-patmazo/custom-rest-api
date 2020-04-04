#!/bin/bash

# Set environment type
ENV=$1

if [ $ENV == "dev" ]
then
    sed -i '/^env /s/=.*$/= "'$ENV'"/' .air.conf && air -d
elif [ $ENV == "qa" ]
then
    sed -i '/^env /s/=.*$/= "'$ENV'"/' .air.conf && air -d
elif [ $ENV == "stg" ]
then
    sed -i '/^env /s/=.*$/= "'$ENV'"/' .air.conf && air -d
elif [ $ENV == "prod" ]
then
    sed -i '/^env /s/=.*$/= "'$ENV'"/' .air.conf && air
fi