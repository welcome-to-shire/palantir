package shire.bcho.palantiroid.notification;

import java.util.ArrayList;

import android.util.Log;

import android.content.Intent;
import android.content.Context;
import shire.bcho.palantiroid.R;
import android.support.v4.app.NotificationCompat;
import android.support.v4.app.NotificationManagerCompat;
import android.support.v4.app.NotificationCompat.BigTextStyle;

/**
 * Notification manager.
 */
public class NotificationManager {

    /**
     * Current in screen notifications.
     */
    private ArrayList<Integer> NOTIFICATION_IDS;

    /**
     * Notification id counter.
     */
    private static int NOTIFICATION_ID_COUNTER = 1000;

    public NotificationManager() {
        NOTIFICATION_IDS = new ArrayList<Integer>();
    }

    /**
     * Mark notification as read.
     */
    public void markAsRead(Context context) {
        Log.d("shire-log", "mark notification as read");

        NotificationManagerCompat nm = NotificationManagerCompat.from(context);
        for (int id : NOTIFICATION_IDS) {
            nm.cancel(id);
        }
    }

    /**
     * Emit a notification.
     */
    public void show(Context context, String title, String content) {
        Log.d("shire-log", "emit notification");

        NotificationCompat.Builder nb = new NotificationCompat.Builder(context)
            .setSmallIcon(R.drawable.ic_launcher)
            .setContentTitle(title)
            .setContentText(content)
            .setStyle(new BigTextStyle().bigText(content));

        NotificationManagerCompat nm = NotificationManagerCompat.from(context);
        nm.notify(getNotificationId(), nb.build());
    }

    /**
     * Get a new notification id.
     */
    public int getNotificationId() {
        int id = NOTIFICATION_ID_COUNTER;
        NOTIFICATION_IDS.add(id);
        NOTIFICATION_ID_COUNTER = NOTIFICATION_ID_COUNTER + 1;

        return id;
    }
}
