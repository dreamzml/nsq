#!/bin/bash
version=`date -d '1 hour ago' '+%y%m%d.%H'`
echo /data/output/test2.$version-* >> /data/cron.log
rm -rf /data/output/test2.$version-*