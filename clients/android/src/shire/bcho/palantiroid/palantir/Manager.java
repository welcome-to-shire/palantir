package shire.bcho.palantiroid.palantir;

import java.net.URL;
import java.net.HttpURLConnection;
import java.io.InputStream;
import java.lang.StringBuilder;
import java.io.InputStreamReader;

import android.util.Log;
import android.widget.Toast;

import com.google.gson.Gson;

import shire.bcho.palantiroid.palantir.model.Message;
import shire.bcho.palantiroid.palantir.model.Error;
import shire.bcho.palantiroid.palantir.PalantirError;

/**
 * Notification service manager.
 */
public class Manager {

    /**
     * API base url.
     */
    private String baseUrl;

    /**
     * Subscribed subject.
     */
    private String subject;

    /**
     * Current api version.
     */
    private final String version = "v1";

    public Manager(String domain, String _subject) {
        // TODO better url construction
        baseUrl = domain + "/api/" + version;
        subject = _subject;
    }

    /**
     * Get a notification.
     */
    public Message GetNotification() throws PalantirError {
        try {
            URL url = new URL(GetFullPath("/" + subject));
            HttpURLConnection conn = (HttpURLConnection) url.openConnection();
            conn.disconnect();
            int status = conn.getResponseCode();

            Gson gson = new Gson();
            InputStreamReader reader;

            switch (status / 100) {
                case 2:
                    reader = new InputStreamReader(conn.getInputStream());
                    Message m = gson.fromJson(reader, Message.class);

                    return m;

                default:
                    reader = new InputStreamReader(conn.getErrorStream());
                    Error err = gson.fromJson(reader, Error.class);
                    Log.d("shire-log", err.reason);

                    if (status == 404) {
                        return null;
                    }

                    throw new PalantirError(err.reason);
            }
        } catch (Exception e) {
            Log.w("shire-log", "error occured");
            throw new PalantirError(e.getMessage());
        }
    }

    /**
     * Generate full access url.
     *
     * Note: ns should start with "/".
     */
    protected String GetFullPath(String ns) {
        return baseUrl + ns;
    }
}
