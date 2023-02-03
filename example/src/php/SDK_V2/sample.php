use Flagship\Flagship;

Flagship::start("your_env_id", "your_api_key");
$visitor = Flagship::newVisitor("your_visitor_id")->build()
$visitor->updateContext("isVip", true)
$visitor->synchronizeModifications();

$displayVipFeature = $visitor->getModification("displayVipFeature", false);
$displayVipFeature = $visitor->getModification("vipFeatureSize", 13);
$displayVipFeature = $visitor->getModification("vipFeatureColor", "red", true);

