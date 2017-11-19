<?php
namespace OCA\RealTimeNotifications\AppInfo;


use OCA\RealTimeNotifications\Authentication\TokenGenerator;
use OCA\RealTimeNotifications\Config;
use OCA\RealTimeNotifications\Controller\AuthController;
use OCA\RealTimeNotifications\Controller\SettingsController;
use OCA\RealTimeNotifications\EventRouterService;
use OCA\RealTimeNotifications\ShareHooks;
use OCP\AppFramework\App;
use OCP\AppFramework\Http\EmptyContentSecurityPolicy;
use OCP\AppFramework\IAppContainer;
use GuzzleHttp\Client as GuzzleClient;

class Application extends App {

	const APP_ID = 'realtime_notifications';

	/** @var Config  */
	private $config;


	/**
	 *
	 * @param array $urlParams
	 */
	public function __construct(array $urlParams = []) {

		parent::__construct(self::APP_ID, $urlParams);

		$container = $this->getContainer();
		$server = $container->getServer();
		$this->config = new Config($server->getConfig());

		$server->getEventDispatcher()->addListener(
			'OCA\Files::loadAdditionalScripts', function() {
				\OCP\Util::addScript(self::APP_ID, 'spop');
				\OCP\Util::addStyle(self::APP_ID, 'spop');
			}
		);

		$server->getContentSecurityPolicyManager()->addDefaultPolicy(
			(new EmptyContentSecurityPolicy())
				->addAllowedConnectDomain($this->config->getBackendHost())
		);

		$container->registerService('EventConfig', function (IAppContainer $c) use ($server) {
			return $this->config;

		});



		$container->registerService('EventAuthTokenGenerator', function (IAppContainer $c) use ($server) {
			return new TokenGenerator($c->query('EventConfig')->getSecret());
		});

		$container->registerService('EventRouterService', function (IAppContainer $c) use ($server) {
			$er = new EventRouterService(
				$c->query('EventConfig'),
				new GuzzleClient(),
				$server->getGroupManager(),
				$c->query('EventAuthTokenGenerator')
			);
			

			ShareHooks::initialize($er);

			return $er;
		});


		$container->query('EventRouterService');

		$container->registerService('SettingsController', function (IAppContainer $c) use ($server) {
			return new SettingsController(
				self::APP_ID,
				$c->query('Request'),
				$c->query('EventConfig')
			);
		});

		$container->registerService('AuthController', function (IAppContainer $c) use ($server) {
			return new AuthController(
				self::APP_ID,
				$c->query('Request'),
				$server->getUserSession(),
				$c->query('EventAuthTokenGenerator')
			);
		});
	}
}
