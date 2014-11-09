package shire.bcho.palantiroid.notification;

import android.util.Log;

import android.os.Handler;
import android.content.Intent;
import android.app.IntentService;

import shire.bcho.palantiroid.palantir.Manager;
import shire.bcho.palantiroid.notification.NotificationManager;
import shire.bcho.palantiroid.palantir.model.Message;

public class NotificationService extends IntentService {

    /**
     * Check every 5 minutes;
     *
     * TODO make it configurable.
     */
    private long checkDurationMillis = 5 * 60 * 1000;

    private Handler timerHandler = new Handler();
    private Runnable task = new Runnable() {
        @Override
        public void run() {
            Message message = messageManager.GetNotification();
            if (message != null) {
                noti.show(getContext(), message.title, message.content);
            }
            timerHandler.postDelayed(this, checkDurationMillis);
        }
    };

    /**
     * Server configuration.
     *
     * TODO Make it configurable.
     */
    private String server = "http://192.168.1.100:9092";
    private String subject = "palantiroid";

    private NotificationManager noti = new NotificationManager();
    private Manager messageManager = new Manager(server, subject);

    public NotificationService() {
        super("Notification Service");
    }

    @Override
    protected void onHandleIntent(Intent intent) {}

    @Override
    public int onStartCommand(Intent intent, int flags, int startId) {
        timerHandler.postDelayed(task, 0);

        return START_STICKY;
    }

    private NotificationService getContext() { return this; }
}
