.PHONY: debug-build debug-install run

PACKAGE=shire.bcho.palantir
MAIN_ACTIVITY=${PACKAGE}.PalantirActivity
TARGET_APK=bin/Palantir-debug.apk

run: debug-install
	adb shell am start -n ${PACKAGE}/${MAIN_ACTIVITY}

debug-install:
	adb uninstall ${PACKAGE}
	adb install -r ${TARGET_APK}

debug-build:
	ant debug
