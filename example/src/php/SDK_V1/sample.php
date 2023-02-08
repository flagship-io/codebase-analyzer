use Flagship\Flagship;

$visitor = Flagship::newVisitor("your_visitor_id")
$visitor->updateContext("isVip", true)
$visitor->synchronizeModifications();

$displayVipFeature = $visitor->getModification("displayVipFeature", false);
$displayVipFeature = $visitor->getModification("vipFeatureSize", 13);
$displayVipFeature = $visitor->getModification("vipFeatureColor", "red", true);
