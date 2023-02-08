public class MyActivity extends Activity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        Flagship.Companion.builder(getApplicationContext(), ENV_ID).start();
    }
}

Flagship.Companion.builder(this.getApplicationContext(), "my_env_id", "my_api_key")

                .withFlagshipMode(Flagship.Mode.BUCKETING)
                .withLogEnabled(Flagship.LogMode.ALL)
                .withVisitorId("my_visitor_id")
                .withReadyCallback(() -> {
                    new Hit.Event(Hit.EventCategory.ACTION_TRACKING, "sdk-android-ready").send();
                    MainJava.this.runOnUiThread(new Runnable() {
                        @Override
                        public void run() {
                            updateView();
                        }
                    });
                    return null;
                })
                .start();

Flagship.Companion.updateContext("vipUser", currentVisitor.vip);
      Flagship.Companion.updateContext("age", currentVisitor.age, () -> {
      MainJava.this.runOnUiThread(this::applyChanges);
      return null;
      });

Flagship.Companion.synchronizeModifications( () -> {
    MainJava.this.runOnUiThread(this::applyChanges);
    return null;
    });

String color = Flagship.Companion.getModification("btnColor", "#FFFFFF");
String color = Flagship.Companion.getModification("backgroundColor", "#000000");
String color = Flagship.Companion.getModification("backgroundSize", 13);
String color = Flagship.Companion.getModification("showBackground", true);

button.setBackgroundColor(Color.parseColor(color));
button.setBackgroundColor(Color.parseColor(color));



// END
