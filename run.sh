
start="run start"
echo "$start"

nohup ./pipe -conf /data/conf/pipe.json > /var/log/server/pipe.log 2>&1 &

end="run end"
echo "$end"