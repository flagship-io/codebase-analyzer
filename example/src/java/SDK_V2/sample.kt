class MyActivity : Activity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        Flagship.builder(applicationContext, ENV_ID).start()
    }
}

Flagship.builder(applicationContext, "my_env_id", "my_api_key")

            .withFlagshipMode(Flagship.Mode.BUCKETING)
            .withLogEnabled(Flagship.LogMode.ALL)
            .withVisitorId("my_visitor_id)
            .withReadyCallback {
                Hit.Event(Hit.EventCategory.ACTION_TRACKING, "sdk-android-ready").send()
                runOnUiThread { update() }
            }
            .start()

val color = Flagship.getModification("btnColor", "#0e5fe3", true)
val size = Flagship.getModification("btnSize", 13, true)
val display = Flagship.getModification("displayBtn", true, true)

button.setBackgroundColor(Color.parseColor(color))

---OR---

Flagship.activateModification("btnColor")
