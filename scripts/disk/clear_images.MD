
apt-get install -y cron

vim etc/crontab

* 1     * * *   root    sh /scripts/disk-monitor.sh