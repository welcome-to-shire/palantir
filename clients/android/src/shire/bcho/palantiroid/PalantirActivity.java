package shire.bcho.palantiroid;

import android.util.Log;

import android.os.Bundle;
import android.app.Activity;
import android.content.Intent;

/**
 * Application main activity.
 */
public class PalantirActivity extends Activity
{
    private final static String BOOT_INTENT = "shire.bcho.palantiroid.BOOT_COMPLETED";

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
    }
}
