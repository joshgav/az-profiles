package com.joshgav.azure.tests;

import com.microsoft.azure.arm.resources.Region;
import com.microsoft.azure.arm.utils.SdkContext;
import com.microsoft.azure.credentials.ApplicationTokenCredentials;
import com.microsoft.azure.AzureEnvironment;

import com.microsoft.azure.management.profile_2018_03_01_hybrid.Azure;

import com.microsoft.azure.management.resources.v2018_02_01.ResourceGroup;
import com.microsoft.azure.management.storage.v2016_01_01.Kind;
import com.microsoft.azure.management.storage.v2016_01_01.Sku;
import com.microsoft.azure.management.storage.v2016_01_01.SkuName;
import com.microsoft.azure.management.storage.v2016_01_01.StorageAccount;
import com.microsoft.azure.management.storage.v2016_01_01.StorageAccountListKeysResult;

import io.github.cdimascio.dotenv.Dotenv;

public class HybridApp {
	private static String clientId;
	private static String tenantId;
	private static String clientSecret;
	private static String subscriptionId;

	protected static String groupName;
	protected static String location;
	protected static String storageAccountName;

	protected static Azure azureStack;


    public static void main(String args[]) throws Exception {
        configureFromEnvironment();

		ApplicationTokenCredentials creds = new ApplicationTokenCredentials(
								clientId,
								tenantId,
								clientSecret,
								AzureEnvironment.AZURE);

        azureStack = Azure.authenticate(creds, subscriptionId);
        createStorageAccount();
        cleanupResources();
    }

    // configureFromEnvironment() sets up context from files and variables in the env
	protected static void configureFromEnvironment() {
		Dotenv dotenv = Dotenv.configure().directory(".").load();

		clientId = dotenv.get("AZURE_CLIENT_ID");
		tenantId = dotenv.get("AZURE_TENANT_ID");
		clientSecret = dotenv.get("AZURE_CLIENT_SECRET");
		subscriptionId = dotenv.get("AZURE_SUBSCRIPTION_ID");

		if ( clientId == null || tenantId == null || clientSecret == null || subscriptionId == null) { throw new IllegalArgumentException("env vars not set");
		}

        // Get 2 random names
        groupName = SdkContext.randomResourceName("rg", 20);
        storageAccountName = SdkContext.randomResourceName("sa", 20);
        location = Region.US_WEST.name();
	}

    public static void createStorageAccount() {
        // Create a resource group
        ResourceGroup resourceGroup = azureStack.resourceGroups().define(groupName)
                .withExistingSubscription()
                .withLocation(location)
                .create();

        // Create a storage account in the resource group
        StorageAccount storageAccount = azureStack.storageAccounts()
                .define(storageAccountName)
                .withRegion(location)
                .withExistingResourceGroup(groupName)
                .withKind(Kind.STORAGE)
                .withSku(new Sku().withName(SkuName.STANDARD_GRS))
                .create();

        System.out.println("Storage account: " + storageAccount.id() + "\nKeys:");

        // List storage account keys
        azureStack.storageAccounts().listKeysAsync(groupName, storageAccountName)
                .flatMapIterable(StorageAccountListKeysResult::keys)
                .doOnNext(key -> System.out.println("\t" +
                                                    key.keyName() +
                                                    ": " +
                                                    key.value()))
                .toBlocking().subscribe();
    }

    protected static void cleanupResources() {
        azureStack.resourceGroups().deleteAsync(groupName).await();
    }

}
