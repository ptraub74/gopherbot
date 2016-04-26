#!/bin/bash -e

# echo.sh - less trivial example shell plugin for gopherbot
set -x
[ -z "$GOPHER_INSTALLDIR" ] && { echo "GOPHER_INSTALLDIR not set" >&2; exit 1; }
source $GOPHER_INSTALLDIR/util/shellLib.sh

HOSTSARR=($*)
for LOOKUP in "${HOSTSARR[@]}"
do
	ERROR=false
	if echo "$LOOKUP" | grep -qP "[a-zA-Z]+"
	then
		HOSTNAME=${LOOKUP##*\|}
		HOSTNAME=${HOSTNAME%\>}
		IPADDR=$(host $HOSTNAME | grep 'has address') || ERROR=true
		IPADDR=${IPADDR##* }
		[ "$ERROR" = "true" ] && IPADDR="(not found)"
	else
		IPADDR=$LOOKUP
		HOSTNAME=$(host $LOOKUP) || ERROR=true
		HOSTNAME=${HOSTNAME##* }
		HOSTNAME=${HOSTNAME%.}
		[ "$ERROR" = "true" ] && HOSTNAME="(not found)"
	fi
	MESSAGE=$(echo -e "${IPADDR}\t${HOSTNAME}\n$MESSAGE")
done

echo "$MESSAGE"
say -f $CHANNEL "$MESSAGE"