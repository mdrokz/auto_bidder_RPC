#include <libnotify/notify.h>
#include <stdio.h>

typedef struct {
    const char *title;
    const char *body;
    const char *link;
    const char *button_name;
    int loop_timeout;
} notificationInfo;

typedef void (*callback)(void *);

static callback _cb;
static void *_user_data;

void register_callback(callback cb, void *user_data);

struct NotificationCallbackInfo {
    const char* link;
    void (*ptr)(const char* link);
};

void notification_callback(NotifyNotification *notification,
              char *action,
              gpointer user_data);

int quit(gpointer user_data);

// void (*ptr)(const char* link)

void send_notification(notificationInfo notification_info);

extern void cb_proxy(void *v);

void _register_callback(void *user_data);