.PHONY: debug-build debug-install run attach-logcat

PACKAGE=shire.bcho.palantiroid
MAIN_ACTIVITY=${PACKAGE}.PalantirActivity
TARGET_APK=bin/Palantir-debug.apk

run: debug-install
	adb shell am start -n ${PACKAGE}/${MAIN_ACTIVITY}

debug-install: debug-build
	adb uninstall ${PACKAGE}
	adb install -r ${TARGET_APK}

debug-build:
	ant debug

attach-logcat:
	adb logcat | grep "/shire-log"
