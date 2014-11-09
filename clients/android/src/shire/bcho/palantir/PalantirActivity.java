package shire.bcho.palantir;

import android.os.Bundle;
import android.app.Activity;
import android.content.Intent;

import shire.bcho.palantir.palantir.Manager;
import shire.bcho.palantir.notification.NotificationManager;
import shire.bcho.palantir.palantir.model.Message;

/**
 * Application main activity.
 */
public class PalantirActivity extends Activity
{
    private final static String BOOT_INTENT = "shire.bcho.palantir.BOOT_COMPLETED";

    @Override
    public void onCreate(Bundle savedInstanceState)
    {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.main);

        // Start notification receiver
        // FIXME receiver should start via `BOOT_COMPLETED`
        Intent intent = new Intent();
        intent.setAction(BOOT_INTENT);
        sendBroadcast(intent);

        Manager m = new Manager("http://192.168.1.100:9092", "palantiroid");
        Message msg = m.GetNotification();
        if (msg != null) {
            NotificationManager.show(this, msg.title, msg.content);
        }
    }
}
