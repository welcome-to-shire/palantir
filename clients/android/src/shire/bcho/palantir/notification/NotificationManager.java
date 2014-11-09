package shire.bcho.palantir.notification;

import android.util.Log;

import shire.bcho.palantir.R;
import android.content.Intent;
import android.content.Context;
import android.support.v4.app.NotificationCompat;
import android.support.v4.app.NotificationManagerCompat;

/**
 * Notification manager.
 */
public class NotificationManager {

    /**
     * Default notification id.
     *
     * TODO support batch notifications.
     */
    private static final int NOTIFICATION_ID = 1000;

    /**
     * Mark notification as read.
     */
    public static void markAsRead(Context context) {
        Log.d("shire-log", "mark notification as read");

        NotificationManagerCompat nm = NotificationManagerCompat.from(context);
        nm.cancel(NOTIFICATION_ID);
    }

    /**
     * Emit a notification.
     */
    public static void show(Context context, String title, String content) {
        Log.w("shire-log", "emit notification");

        NotificationCompat.Builder nb = new NotificationCompat.Builder(context)
            .setSmallIcon(R.drawable.ic_launcher)
            .setContentTitle(title)
            .setContentText(content);

        NotificationManagerCompat nm = NotificationManagerCompat.from(context);
        nm.notify(NOTIFICATION_ID, nb.build());
    }
}
