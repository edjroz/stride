#!/bin/sh

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

echo 'Initializing Hermes...'

HERMES_CONFIG_FILE="${STATE}/hermes_config.toml"
cp ${SCRIPT_DIR}/hermes_config_template.toml $HERMES_CONFIG_FILE
sed -i -E "s|STRIDE_ENDPOINT|$STRIDE_MAIN_ENDPOINT|g" $HERMES_CONFIG_FILE
sed -i -E "s|GAIA_ENDPOINT|$GAIA_MAIN_ENDPOINT|g" $HERMES_CONFIG_FILE

HERMES_STRIDE_MNEMONIC=$(GET_MNEMONIC $HERMES_STRIDE_ACCT)
HERMES_GAIA_MNEMONIC=$(GET_MNEMONIC $HERMES_GAIA_ACCT)

HERMES_STARTUP_FILE="${STATE}/hermes_startup.sh"
cp ${SCRIPT_DIR}/hermes_startup_template.sh $HERMES_STARTUP_FILE
sed -i -E "s|STRIDE_ENDPOINT|$STRIDE_MAIN_ENDPOINT|g" $HERMES_STARTUP_FILE
sed -i -E "s|HERMES_STRIDE_MNEMONIC|$HERMES_STRIDE_MNEMONIC|g" $HERMES_STARTUP_FILE
sed -i -E "s|HERMES_GAIA_MNEMONIC|$HERMES_GAIA_MNEMONIC|g" $HERMES_STARTUP_FILE

rm -f "${HERMES_CONFIG_FILE}-e"
rm -f "${HERMES_STARTUP_FILE}-e"