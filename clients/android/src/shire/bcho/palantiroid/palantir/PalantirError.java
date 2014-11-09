package shire.bcho.palantiroid.palantir;

/**
 * Basic exception.
 */
public class PalantirError extends Exception {

    public String reason;

    public PalantirError() {
        reason = "unknown error";
    }

    public PalantirError(String _reason) {
        reason = _reason;
    }

    public String getMessage() { return reason; }
}
