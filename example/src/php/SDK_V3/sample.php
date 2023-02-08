use Flagship\Flagship;

Flagship::start("your_env_id", "your_api_key");
$visitor = Flagship::newVisitor("your_visitor_id")->build();
$visitor->updateContext("isVip", true);
$visitor->fetchFlags();

$flag = $visitor->getFlag("displayVipFeature", false);
$flag = $visitor->getFlag("vipFeatureSize", 13);
$flag = $visitor->getFlag("vipFeatureColor", "red");
