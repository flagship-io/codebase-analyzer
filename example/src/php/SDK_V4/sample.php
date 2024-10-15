use Flagship\Flagship;

// Step 1: start the SDK
Flagship::start("<ENV_ID>", "<API_KEY>");

        //Step 2: Create a visitor
        $visitor = Flagship::newVisitor("<VISITOR_ID>", true)
            ->setContext(["isVip" => true])
            ->build();

            //Step 3: Fetch flag from the Flagship platform
            $visitor->fetchFlags();

            //Step 4: Retrieves a flag named "displayVipFeature"
            $flag = $visitor->getFlag("displayVipFeature");

            //Step 5: Returns the flag value ,or if the flag does not exist, it returns the default value "false"
            echo "flag value:". $flag->getValue(false);

            $flag1 = $visitor->getFlag("vipFeatureSize")->getValue(15);

            $flag1 = $visitor->getFlag("vipFeatureColor")->getValue("red");

            //Step 6: Batch all the collected hits and send them
            Flagship::close();