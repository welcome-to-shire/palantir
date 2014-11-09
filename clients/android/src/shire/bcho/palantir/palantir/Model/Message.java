package shire.bcho.palantir.palantir.model;

import com.google.gson.annotations.SerializedName;

public class Message {

    @SerializedName("title")
    public String title;

    @SerializedName("content")
    public String content;

    @SerializedName("created_at")
    public String createdAt;
}
