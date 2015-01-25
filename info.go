package slack

import (
	"fmt"
	"time"
)

// XXX: Need to implement
type UserPrefs struct {
	// "highlight_words":"",
	// "user_colors":"",
	// "color_names_in_list":true,
	// "growls_enabled":true,
	// "tz":"Europe\/London",
	// "push_dm_alert":true,
	// "push_mention_alert":true,
	// "push_everything":true,
	// "push_idle_wait":2,
	// "push_sound":"b2.mp3",
	// "push_loud_channels":"",
	// "push_mention_channels":"",
	// "push_loud_channels_set":"",
	// "email_alerts":"instant",
	// "email_alerts_sleep_until":0,
	// "email_misc":false,
	// "email_weekly":true,
	// "welcome_message_hidden":false,
	// "all_channels_loud":true,
	// "loud_channels":"",
	// "never_channels":"",
	// "loud_channels_set":"",
	// "show_member_presence":true,
	// "search_sort":"timestamp",
	// "expand_inline_imgs":true,
	// "expand_internal_inline_imgs":true,
	// "expand_snippets":false,
	// "posts_formatting_guide":true,
	// "seen_welcome_2":true,
	// "seen_ssb_prompt":false,
	// "search_only_my_channels":false,
	// "emoji_mode":"default",
	// "has_invited":true,
	// "has_uploaded":false,
	// "has_created_channel":true,
	// "search_exclude_channels":"",
	// "messages_theme":"default",
	// "webapp_spellcheck":true,
	// "no_joined_overlays":false,
	// "no_created_overlays":true,
	// "dropbox_enabled":false,
	// "seen_user_menu_tip_card":true,
	// "seen_team_menu_tip_card":true,
	// "seen_channel_menu_tip_card":true,
	// "seen_message_input_tip_card":true,
	// "seen_channels_tip_card":true,
	// "seen_domain_invite_reminder":false,
	// "seen_member_invite_reminder":false,
	// "seen_flexpane_tip_card":true,
	// "seen_search_input_tip_card":true,
	// "mute_sounds":false,
	// "arrow_history":false,
	// "tab_ui_return_selects":true,
	// "obey_inline_img_limit":true,
	// "new_msg_snd":"knock_brush.mp3",
	// "collapsible":false,
	// "collapsible_by_click":true,
	// "require_at":false,
	// "mac_ssb_bounce":"",
	// "mac_ssb_bullet":true,
	// "win_ssb_bullet":true,
	// "expand_non_media_attachments":true,
	// "show_typing":true,
	// "pagekeys_handled":true,
	// "last_snippet_type":"",
	// "display_real_names_override":0,
	// "time24":false,
	// "enter_is_special_in_tbt":false,
	// "graphic_emoticons":false,
	// "convert_emoticons":true,
	// "autoplay_chat_sounds":true,
	// "ss_emojis":true,
	// "sidebar_behavior":"",
	// "mark_msgs_read_immediately":true,
	// "start_scroll_at_oldest":true,
	// "snippet_editor_wrap_long_lines":false,
	// "ls_disabled":false,
	// "sidebar_theme":"default",
	// "sidebar_theme_custom_values":"",
	// "f_key_search":false,
	// "k_key_omnibox":true,
	// "speak_growls":false,
	// "mac_speak_voice":"com.apple.speech.synthesis.voice.Alex",
	// "mac_speak_speed":250,
	// "comma_key_prefs":false,
	// "at_channel_suppressed_channels":"",
	// "push_at_channel_suppressed_channels":"",
	// "prompted_for_email_disabling":false,
	// "full_text_extracts":false,
	// "no_text_in_notifications":false,
	// "muted_channels":"",
	// "no_macssb1_banner":false,
	// "privacy_policy_seen":true,
	// "search_exclude_bots":false,
	// "fuzzy_matching":false
}

type UserDetails struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	Created        JSONTime  `json:"created"`
	ManualPresence string    `json:"manual_presence"`
	Prefs          UserPrefs `json:"prefs"`
}

type JSONTime int64

func (t JSONTime) String() string {
	tm := time.Unix(int64(t), 0)
	return fmt.Sprintf("\"%s\"", tm.Format("Mon Jan _2"))
}

type Team struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Domain string `json:"name"`
}

type Info struct {
	Url      string      `json:"url,omitempty"`
	User     UserDetails `json:"self,omitempty"`
	Team     Team        `json:"team,omitempty"`
	Users    []User      `json:"users,omitempty"`
	Channels []Channel   `json:"channels,omitempty"`
}

type infoResponseFull struct {
	Info
	SlackResponse
}

func (info Info) GetUserById(id string) *User {
	for _, user := range info.Users {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

func (info Info) GetChannelById(id string) *Channel {
	for _, channel := range info.Channels {
		if channel.Id == id {
			return &channel
		}
	}
	return nil
}