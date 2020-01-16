#!/bin/bash

zoneurl=https://data.iana.org/time-zones/tzdb-2019c/zone.tab
tmpfile="./tz.tmp"
jsonfile="./tz.json"
: > "${jsonfile}"

curl -s -o "${tmpfile}" "${zoneurl}"

echo "[" > "${jsonfile}"

main(){
    egrep -v \# "${tmpfile}" | while read line;do
        #echo ${line}
        code=$(echo "${line}" | awk '{print $1}')
        coordinates=$(echo "${line}" | awk '{print $2}')
        tz=$(echo "${line}" | awk '{print $3}')
        region=$(echo "${line}" | awk '{print $3}' | awk -F \/ '{print $1}')
        zone_1=$(echo "${line}" | awk '{print $3}' | awk -F \/ '{print $2}')
        zone_2=$(echo "${line}" | awk '{print $3}' | awk -F \/ '{print $3}')
        echo -n -e "\t"
        echo -n "{"
        echo -n "\"region\":\""${region}"\","
        echo -n "\"zone_1\":\""${zone_1}"\","
        echo -n "\"zone_2\":\""${zone_2}"\","
        echo -n "\"code\":\""${code}"\","
        echo -n "\"tz\":\""${tz}"\","
        echo -n "\"coordinates\":\""${coordinates}"\""
        echo "},"
    done
}

main >> "${jsonfile}"

lastline=$(tail -1 "${jsonfile}")
lastline_sed=$(tail -1 "${jsonfile}" | sed -e "s|,$||")
sed -i "" -e "s|${lastline}|${lastline_sed}|" "${jsonfile}"

echo "]" >> "${jsonfile}"

#cat "${jsonfile}"
rm "${tmpfile}"