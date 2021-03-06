#!/bin/bash -e

# whoami.sh - shell plugin example that retrieves user attributes
[ -z "$GOPHER_INSTALLDIR" ] && { echo "GOPHER_INSTALLDIR not set" >&2; exit 1; }
source $GOPHER_INSTALLDIR/util/shellLib.sh

if [ "$1" != "whoami" ]
then
	exit 0
fi
shift

USERFULLNAME=$(GetSenderAttribute fullName)
USEREMAIL=$(GetSenderAttribute email)
Reply "You're $USERFULLNAME, $USEREMAIL"
