#!/bin/bash

until /home/dinamyc/Code/go/src/Maghaze_Bot/Maghaze_Bot ;  do
        echo "Server 'Mmaghaze_Bot' crashed with exit code $?. Respawning.." >&2
        sleep 1
done

