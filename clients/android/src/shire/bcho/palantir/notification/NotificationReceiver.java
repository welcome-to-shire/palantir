package shire.bcho.palantir.notification;

import android.util.Log;

import android.content.Intent;
import android.content.Context;
import android.content.BroadcastReceiver;

/**
 * Boot up notification check service.
 */
public class NotificationReceiver extends BroadcastReceiver {

    @Override
    public void onReceive(Context context, Intent intent) {
        Log.w("shire-log", "receiver booted up");
    }
}
